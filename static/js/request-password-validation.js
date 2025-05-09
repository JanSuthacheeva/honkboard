const input = document.getElementById('code');
input.addEventListener('input', () => {
    input.value = input.value.replace(/[^0-9]/g, '');
});
