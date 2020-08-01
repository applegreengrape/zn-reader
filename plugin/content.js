// Let the background stript know if we detect Chinese.
if (document.body.innerText.match(/[\u3400-\u9FBF]/)) {
    let msg = { "direction": "from_content_script", "event": "detected_zn" };
    console.log("zn-reader", "content_script", "sending", msg);
    browser.runtime.sendMessage(msg);
}

// ExtractZN starts at node 'n' and recursively traverses the DOM returning
// an array containing all strings of Chinese characters. The following node
// types are skipped: SCRIPT, STYLE.
function extractZN(n) {
    let strings = [];

    if (n.nodeType == 3) {
        return znStrings(n.textContent);
    }

    if (n.nodeType != 1) {
        return strings;
    }

    const skipElems = ["SCRIPT", "STYLE"]
    if (skipElems.includes(n.tagName)) {
        return strings;
    }

    n.childNodes.forEach((zn, index) => {
        strings = strings.concat(extractZN(zn));
    });

    return strings
}

// znStrings takes a string and returns an array containing all substrings of
// Chinese characters. For example: BBC英伦网2020年 will return [英伦网, 年].
function znStrings(t) {
    let strings = [];
    for (let i = 0; i < t.length; i++) {
        if (!isZN(t[i])) {
            continue;
        }

        for (let j = i; j < t.length; j++) {
            if (isZN(t[j])) {
                continue;
            }

            if (i != j) {
                strings.push(t.slice(i, j));
            }

            i = j + 1;
        }
    }
    return strings;
}

// IsZN tests to see if a character is Chinese.
// TODO: validate these character codes. They work for now.
function isZN(c) {
    c = c.charCodeAt(0);
    return 19968 <= c && 40879 >= c;
}

// ScanDoc returns a promise that contains Chinese strings present in the
// current document body.
async function scanDoc() {
    let zn = await extractZN(document.body)
    let msg = {
        "strings": zn,
        "hash": null
    };
    return msg;
}

// Wait until the user asks for Chinese markup by clicking on the page_action 
// before extracting any page content.
browser.runtime.onMessage.addListener(request => {
    console.log("zn-reader", "content_script", "received", request);
    if (request.direction == 'from_page_action' && request.event == 'request_zn') {
        return scanDoc();
    }

    if (request.direction == 'from_page_action' && request.event == 'return_zn') {
        console.log("badaboom");
    }

    return null;
});