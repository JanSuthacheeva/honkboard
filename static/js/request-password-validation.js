document.addEventListener('alpine:init', () => {
     Alpine.data('verificationCode', () => {
      return {
        code: ['', '', '', '', '', ''],
        joinedCode() {
            return this.code.join("");
        },
        init() {
          // Wait for the DOM to fully render
          this.$nextTick(() => {
            this.$el.querySelectorAll('.code-input').forEach((input, index) => {
              input.addEventListener('input', (event) => this.updateCode(event, index));
              input.addEventListener('keydown', (event) => {
                if (event.key === 'Backspace') {
                  this.handleBackspace(event, index);
                }
              });
            });
          });
        },
        updateCode(event, index) {
          const value = event.target.value;
          if (/^\d$/.test(value)) {
            this.code[index] = value;
            // Move to the next input if not the last one
            if (index < this.code.length - 1) {
              event.target.nextElementSibling?.focus();
            }
          } else {
            // Clear invalid input
            this.code[index] = '';
            event.target.value = '';
          }
        },
        handleBackspace(event, index) {
          if (!this.code[index] && index > 0) {
            // Move to the previous input if empty
            event.target.previousElementSibling?.focus();
          }
        }
      };
    });
});
