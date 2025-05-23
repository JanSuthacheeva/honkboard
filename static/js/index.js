document.addEventListener('alpine:init', () => {
	Alpine.data('modals', () => {
		return {
			showDeleteAllModal: false,
			closeDeleteAllModal() {
				this.showDeleteAllModal = false;
				Array.from(document.getElementsByClassName('list-switcher')).forEach((el) => {
					el.classList.remove("pointer-events-none");
					el.classList.remove("blur");
				});
				document.getElementsByTagName('footer')[0].classList.remove('pointer-events-none');
			},
			openDeleteAllModal() {
				this.showDeleteAllModal = true;
				Array.from(document.getElementsByClassName('list-switcher')).forEach((el) => {
					el.classList.add("pointer-events-none");
					el.classList.add("blur");
				});
				document.getElementsByTagName('footer')[0].classList.add('pointer-events-none');
			},
			get pointerClass() {
				return this.showDeleteAllModal ? 'pointer-events-none blur' : '';
			},
			showDeleteTodoModal: false,
			closeDeleteTodoModal() {
				this.showDeleteTodoModal = false;
			},
			openDeleteTodoModal() {
				this.showDeleteTodoModal = true;
			}
		}
	});
});

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
