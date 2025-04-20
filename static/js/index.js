document.addEventListener("htmx:afterRequest", function(event) {
	event.target.blur();
	if (event.detail.successful) {
		if (event.target.tagName === "FORM") {
				event.target.reset();
		}
	}
});

Array.from(document.getElementsByClassName('list-switcher')).forEach((el) => {
	el.addEventListener('click', () => {
		const targetElement = document.querySelector('#todo-list');
		targetElement.classList.add('swipe-enabled');

	// Optional: Remove the class after a delay (e.g., after the animation ends)
		setTimeout(() => {
			targetElement.classList.remove('swipe-enabled');
		}, 400); // Adjust the timeout to match the animation duration
		});
});
