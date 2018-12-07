var images = [
	"1.png",
	"2.png",
	"3.jpg",
	"4.jpg",
	"5.png",
	"6.jpg",
];
var nextImage = 1;

function setImg(id, val) {
	var pic = document.getElementById(id);
	pic.src = "/Gooeyboi/pictures/"+val;
}

function swapImage() {
	if (nextImage == images.length) {
		nextImage = 0;
		setImg("pic1", images[nextImage]);
	} else {
		setImg("pic1", images[nextImage]);
	}
	nextImage++;
}

function google(thing) {
	w = window.open("https://www.google.com/search?q="+encodeURIComponent(thing));
}
