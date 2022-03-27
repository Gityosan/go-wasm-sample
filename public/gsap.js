window.onload = ()=>{
	gsap.from("#title", {x: 1000, duration: 1});
	gsap.from("#result", {scale:0, duration: 2, ease:"expo",rotation: 1440});
	gsap.from(".calculator>div>button", {scale:0, duration: 2, ease:"expo",rotation: 1440})

}