var aboutEl = document.querySelector("#about");
var projectEl = document.querySelector("#projects");
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
        case "p": {
            projectEl.click();
            break;
        }
        default:
            return;
    }
})