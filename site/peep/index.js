var images = ["1.png", "2.png", "3.png"];
var nextImage = 1;

// For the cookie button
var cookieClicks = 0;
var cookieTexts = [
	"please don't i really like cookies",
	"okay fine if you really want them",
	"SIKE ALL MINE NERD",
];

var memeRevealed = false;
var memes = [
	// The youtube id for meme videos
	// the part after watch?v=
	"ocwnns57cYQ", // The mestryos life of develports
	"74gx0v5ZJzw", // trump havana
	"BivAv_hBJVY",
	"KjdE7zz3BqA",
	"GVN17U3Vg34", // trump do you want to build a wall
	"NCcvPXDxA2k",
	"M2qroMuIluI",
	"3nx7_G5R0oA",
	"on9UFsxdC0Q", // trump new rules
	//"bmhbqKT7ONo", // Obama shape of you
];

// Utility functions
function sleep(ms) {
	return new Promise(resolve => setTimeout(resolve, ms * 1000));
}

function google(thing) {
	w = window.open("https://www.google.com/search?q="+encodeURIComponent(thing));
}

function setImg(id, val) {
	var pic = document.getElementById(id);
	pic.src = "pictures/"+val;
}

// Swaps to the next image in the images list.
function swapImage() {
	if (nextImage == images.length) {
		nextImage = 0;
		setImg("pic1", images[nextImage]);
	} else {
		setImg("pic1", images[nextImage]);
	}
	nextImage++;
}

// This gets executed when the user clicks the free cookies button.
function cookies() {
	// I'd like to have this global except it won't work because
	// cookiebutton does not exist yet ):
	var cookieButton = document.getElementById("cookieButton");

	if (cookieClicks == cookieTexts.length - 1) {
		window.open("/peep/cookie");
		cookieButton.innerText = cookieTexts[cookieTexts.length - 1];
	} else {
		cookieButton.innerText = cookieTexts[cookieClicks];
		cookieClicks++;
	}
}

// Changes the revealed meme video
function changeVideo(videoID) {
	video = document.getElementById("hiddenVideo"); // The hidden div
	vid   = document.getElementById("video"); // the video
	if (memeRevealed == false) {
		video.style.display = "block";
		memeRevealed = true;
	}

	// cuz i hate the flicker when its one video
	if (vid.src == "https://www.youtube.com/embed/"+videoID+"?ecver=1") {
		return false
	}
	// pro string manipulation B
	vid.src = "https://www.youtube.com/embed/"+videoID+"?ecver=1"
	return true
}

// Change the video to a random meme
function meme() {
	while (true) {
		choice = Math.floor(Math.random() * memes.length);
		if (changeVideo(memes[choice]) == true) {
			break
		}
	}
}
