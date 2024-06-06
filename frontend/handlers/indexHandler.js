let backgroundCenterX = 3000 / 2;
let backgroundCenterY = 2000 / 2;
let scrolloffset = 0;

addEventListener("mousemove", function(e) {
    let x = e.clientX;
    let y = e.clientY;
    
    this.document.body.style.backgroundPositionX = backgroundCenterX + "px ";
    this.document.body.style.backgroundPositionY = backgroundCenterY + "px ";

    this.document.body.style.backgroundPositionX = x / 10 - backgroundCenterX / 2 + "px ";
    this.document.body.style.backgroundPositionY = y / 10 - backgroundCenterY / 2 + "px ";
})

