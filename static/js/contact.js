const contactForm = document.getElementById('contactForm');
const contactFormMessage = document.querySelector('.form-message');

const showMessage = (type, message) => {
  contactFormMessage.textContent = message;
  contactFormMessage.className = `form-message ${type}`;

  setTimeout(() => {
    contactFormMessage.className = 'form-message';
    contactFormMessage.textContent = '';
  }, 5000);
};

contactForm.addEventListener('submit', async (e) => {
  e.preventDefault();

  const formData = new FormData(contactForm);
  const data = Object.fromEntries(formData);

  try {
    const response = await fetch('/api/contact', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    const result = await response.json();
    if (!response.ok) {
      throw new Error(result.error || 'Failed to send message');
    }
    showMessage('Success', result.message);
    contactForm.reset();
  } catch (error) {
    showMessage('Error', error.message);
  }
});
