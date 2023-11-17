var aboutEl = document.querySelector("#about");
var homeLinkEl = document.querySelector("#home-lnk");

document.addEventListener("keyup", function(event) {

    switch (event.key) {
        case "a": {
            aboutEl.click();
            break;
        }
        case "h": {
            homeLinkEl.click();
            break;
        }
        default:
            console.error("shouldn't happen");

    }
})