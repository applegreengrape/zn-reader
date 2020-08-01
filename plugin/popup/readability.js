browser.runtime.sendMessage({ "popup": "shown" });

browser.tabs.query({
    currentWindow: true,
    active: true
}).then(requestZN).catch(onError);

function requestZN(tabs) {
    let tab = tabs[0];
    let message = {
        "direction": "from_page_action",
        "event": "request_zn"
    };

    console.log("zn-reader", "page_action", "sending", message);
    browser.tabs.sendMessage(tab.id, message).then(response => {
        console.log("zn-reader", "page_action", "received", response);

        postData('https://httpbin.org/anything', response).then(data => {
            console.log("zn-reader", "page_action", "received", data);
            displayReadability(tab, data);
        });

    }).catch(onError);
}

function displayReadability(tab, data) {
    browser.tabs.sendMessage(tab.id, {
        "direction": "from_page_action",
        "event": "return_zn",
        "data": data
    });

    document.body.innerHTML = "<strong>Readability</strong>: 57%<p id='data'></p>";

    data = JSON.parse(data.data);
    document.getElementById('data').innerText = data.strings;
}

function onError(error) {
    console.log("zn-reader", "page_action", "error", `${error}`);
}

async function postData(url = '', data = {}) {
    console.log("✅", "postData");
    const response = await fetch(url, {
        method: 'POST',
        mode: 'cors',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
    return response.json();
}