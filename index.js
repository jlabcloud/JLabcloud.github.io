const dropdownTriggers = document.querySelectorAll('[id^="dropdown-trigger"]');
const dropdownMenus = document.querySelectorAll('[id^="dropdown-menu"]');

dropdownTriggers.forEach((trigger, index) => {
    const menu = dropdownMenus[index];

    trigger.addEventListener('click', (event) => {
        event.preventDefault();
        const isVisible = menu.style.display === 'block';
        dropdownMenus.forEach((m) => (m.style.display = 'none')); // Close other menus
        menu.style.display = isVisible ? 'none' : 'block';
    });
});

document.addEventListener('click', (event) => {
    dropdownTriggers.forEach((trigger, index) => {
        const menu = dropdownMenus[index];
        if (!trigger.contains(event.target) && !menu.contains(event.target)) {
            menu.style.display = 'none';
        }
    });
});

// Cart dropdown logic
const cartTrigger = document.getElementById('cart-trigger');
const cartMenu = document.getElementById('cart-menu');

cartTrigger.addEventListener('click', (event) => {
    event.preventDefault();
    const isVisible = cartMenu.style.display === 'block';
    cartMenu.style.display = isVisible ? 'none' : 'block';
});

document.addEventListener('click', (event) => {
    if (!cartTrigger.contains(event.target) && !cartMenu.contains(event.target)) {
        cartMenu.style.display = 'none';
    }
});