const modeToggler = document.querySelector('.navbar .mode-toggler');

const toggleTheme = () => {
  modeToggler.classList.toggle('dark-theme');
  document.body.classList.toggle('dark-theme');
};

modeToggler.addEventListener('click', toggleTheme);

const navLinks = document.querySelectorAll('.nav-links .link');

navLinks.forEach((link) => {
  link.addEventListener('click', () => {
    navLinks.forEach((nav) => {
      nav.classList.remove('active');
    });

    link.classList.add('active');
  });
});

const hamburgerMenu = document.querySelector('.hamburger-menu');

hamburgerMenu.addEventListener('click', () => {
  hamburgerMenu.classList.toggle('active');
});

navLinks.forEach((link) => {
  link.addEventListener('click', () => {
    hamburgerMenu.classList.remove('active');
  });
});
