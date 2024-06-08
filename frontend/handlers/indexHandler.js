let backgroundCenterX = 3000 / 2;
let backgroundCenterY = 2000 / 2;

addEventListener("scroll", function(e)
{
    let screenY = this.window.scrollY;
    const h1Text = document.getElementsByClassName("startingText");
    const loginWnd = document.getElementsByClassName("loginWindow");
    console.log(screenY);
    if (screenY >= 400)
        {
            this.setTimeout(function() {
                h1Text[0].style.opacity = 0;
                h1Text[0].style.transition = "all 0.3s ease-in-out";
            }, 10);
        } else {
            this.setTimeout(function() {
                h1Text[0].style.opacity = 1;
                h1Text[0].style.transition = "all 0.3s ease-in-out";
            }, 10);
        }
    
    if (screenY >= 417)
        {

            this.setTimeout(function() {
                this.disabled = false;
                loginWnd[0].style.opacity = 1;
                loginWnd[0].style.transition = "all 0.3s ease-in-out";
            }, 30);
        } else {
            this.disabled = true;
            this.setTimeout(function() {
                loginWnd[0].style.opacity = 0;
                loginWnd[0].style.transition = "all 0.3s ease-in-out";
            }, 30);
        }

})
addEventListener("mousemove", function(e) {
    let x = e.clientX;
    let y = e.clientY;
    
    this.document.body.style.backgroundPositionX = backgroundCenterX + "px ";
    this.document.body.style.backgroundPositionY = backgroundCenterY + "px ";

    this.document.body.style.backgroundPositionX = x / 10 - backgroundCenterX / 2 + "px ";
    this.document.body.style.backgroundPositionY = y / 10 - backgroundCenterY / 2 + "px ";
})

let btnLogin = document.getElementById("login");
btnLogin.addEventListener("mouseenter", function(e)
{
    const usernameData = document.getElementById("usrname");
    if (usernameData.value == "")
        {
            btnLogin.disabled = true;
            btnLogin.style.backgroundColor = "rgba(117, 38, 38, 0.548)"
            btnLogin.style.transition = "all 0.3s ease-in-out";
        } else {
            btnLogin.style.backgroundColor = "white";
            btnLogin.style.transition = "all 0.3s ease-in-out";
        }
})
btnLogin.addEventListener("mouseleave", function(e) {
    const usernameData = document.getElementById("usrname");
    btnLogin.style.backgroundColor = "white";
    btnLogin.disabled = false;
})