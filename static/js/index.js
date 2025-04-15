document.addEventListener("htmx:afterRequest", function(event) {
	event.target.blur();
	if (event.detail.successful) {
		if (event.target.tagName === "FORM") {
				event.target.reset();
		}
	}
});
