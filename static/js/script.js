const modeToggler = document.querySelectorAll('.change-mode');

// Apply theme based on the value in localStorage
const applyTheme = (theme) => {
  if (theme === 'dark') {
    document.body.classList.add('dark-theme');
    modeToggler.forEach((toggler) => {
      toggler.classList.add('dark-theme');
    });
  } else {
    document.body.classList.remove('dark-theme');
    modeToggler.forEach((toggler) => {
      toggler.classList.remove('dark-theme');
    });
  }
};

const toggleTheme = () => {
  const currentTheme = document.body.classList.contains('dark-theme')
    ? 'dark'
    : 'light';
  const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
  applyTheme(newTheme);
  localStorage.setItem('theme', newTheme);
};
const savedTheme = localStorage.getItem('theme') || 'light';
applyTheme(savedTheme);

modeToggler.forEach((toggler) => {
  toggler.addEventListener('click', toggleTheme);
});

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
