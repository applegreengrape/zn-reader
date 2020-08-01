browser.runtime.onMessage.addListener(notify);

function notify(request, sender) {
    console.log("zn-reader", "background", "received", request);
    if (request.direction == 'from_content_script' && request.event == 'detected_zn') {
        browser.pageAction.show(sender.tab.id);
    }
}
