var aboutEl = document.querySelector("#about");
var projectEl = document.querySelector("#projects");
var homeLinkEl = document.querySelector("#home-lnk");
var blogEl = document.querySelector("#blog");
var projectAbsEls = document.querySelectorAll(".project-abstract");
var screenSize = window.innerWidth;
var homeVidEl = document.querySelector("#home-vid");

if (!localStorage.getItem("hotkeys") && window.innerWidth > 800 ) {
    alert("Hotkeys: a: about, p: project, b: blog, h: home");
    localStorage.setItem("hotkeys", "true");
}

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
        case "b": {
            blogEl.click();
            break;
        }
        default:
            return;
    }
})

function truncateText(sentence) {
    if (sentence.length >= 200) {
        var newString = sentence.slice(0, 200);
        return newString + "...";
    }
}

if (window.innerWidth > 600) {
    homeVidEl.setAttribute("width", "400px");
} 

for (let i=0;i<projectAbsEls.length;i++) {
    projectAbsEls[i].textContent = truncateText(projectAbsEls[i].textContent);
}