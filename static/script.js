const modeToggler = document.querySelector('.navbar .mode-toggler');

const toggleTheme = () => {
  modeToggler.classList.toggle('dark-theme');
  document.body.classList.toggle('dark-theme');
};

modeToggler.addEventListener('click', toggleTheme);

const navs = document.querySelectorAll('.navbar .navs .nav-links .link');

navs.forEach((link) => {
  link.addEventListener('click', () => {
    navs.forEach((nav) => {
      nav.classList.remove('active');
    });

    link.classList.add('active');
  });
});
