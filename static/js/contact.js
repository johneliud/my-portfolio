class Notification {
  constructor() {
    this.timeout = null;
    this.element = null;
  }

  createNotificationElement(type, message) {
    const notification = document.createElement('div');
    notification.className = `notification ${type.toLowerCase()}`;

    notification.innerHTML = `
      ${
        type.toLowerCase() === 'success'
          ? `<svg class="notification-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
             <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"></path>
           </svg>`
          : `<svg class="notification-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
             <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"></path>
           </svg>`
      }
      <span class="notification-message">${message}</span>
      <button class="notification-close" aria-label="Close notification">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
    `;

    return notification;
  }

  show(type, message, duration = 5000) {
    if (this.element) {
      this.element.remove();
      clearTimeout(this.timeout);
    }

    this.element = this.createNotificationElement(type, message);
    document.body.appendChild(this.element);

    const closeButton = this.element.querySelector('.notification-close');
    closeButton.addEventListener('click', () => this.hide());

    requestAnimationFrame(() => {
      this.element.classList.add('show');
    });

    this.timeout = setTimeout(() => {
      this.hide();
    }, duration);
  }

  hide() {
    if (this.element) {
      this.element.classList.remove('show');
      setTimeout(() => {
        this.element.remove();
        this.element = null;
      }, 300);
    }
  }
}

// Initialize notification system
const notification = new Notification();

// Form submission handler
document.getElementById('contactForm').addEventListener('submit', async (e) => {
  e.preventDefault();

  const formData = new FormData(e.target);
  const data = Object.fromEntries(formData);

  try {
    const response = await fetch('/contact', {
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

    notification.show('success', result.message);
    e.target.reset();
  } catch (error) {
    notification.show('error', error.message);
  }
});
