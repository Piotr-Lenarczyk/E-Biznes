const FRONTEND_URL = 'http://172.23.224.32:3000';

describe('template spec', () => {
  it('passes', () => {
    cy.visit('https://example.cypress.io')
  })
});

describe('Products Page', () => {
  beforeEach(() => {
    cy.visit(`${FRONTEND_URL}/`);
  });

  it('should display the Products header', () => {
    cy.get('h2').contains('Products').should('be.visible');
  });

  it('should list all products', () => {
    cy.get('ul > li').should('have.length.greaterThan', 0);
  });

  it('should display product names', () => {
    cy.get('ul > li').each(($el) => {
      cy.wrap($el).should('not.be.empty');
    });
  });
});

describe('Create Cart Page', () => {
  beforeEach(() => {
    cy.visit(`${FRONTEND_URL}/create-cart`);
  });

  it('should display the Create Cart header', () => {
    cy.get('h2').contains('Create Cart').should('be.visible');
  });

  it('should display a loading message if products are not loaded', () => {
    cy.get('p').contains('Loading products...').should('be.visible');
  });

  it('should list all available products with checkboxes', () => {
    cy.get('input[type="checkbox"]').should('have.length.greaterThan', 0);
  });

  it('should allow selecting products', () => {
    cy.get('input[type="checkbox"]').first().check().should('be.checked');
  });

  it('should allow deselecting products', () => {
    cy.get('input[type="checkbox"]').first().check().uncheck().should('not.be.checked');
  });

  it('should submit the form and create a cart', () => {
    cy.get('input[type="checkbox"]').first().check();
    cy.get('button[type="submit"]').click();
    cy.url().should('include', '/cart/');
  });
});

describe('Cart Page', () => {
  beforeEach(() => {
    cy.visit(`${FRONTEND_URL}/cart/2`);
  });

  it('should display the Cart header with ID', () => {
    cy.get('h2').contains('Cart #2').should('be.visible');
  });

  it('should list products in the cart', () => {
    cy.get('ul > li').should('have.length.greaterThan', 0);
  });

  it('should display all available products with checkboxes', () => {
    cy.get('input[type="checkbox"]').should('have.length.greaterThan', 0);
  });

  it('should allow selecting additional products', () => {
    cy.get('input[type="checkbox"]').first().check().should('be.checked');
  });

  it('should allow deselecting products', () => {
    cy.get('input[type="checkbox"]').first().check().uncheck().should('not.be.checked');
  });

  it('should update the cart when clicking the update button', () => {
    cy.get('input[type="checkbox"]').first().check();
    cy.get('button').contains('Update Cart').click();
    cy.get('ul > li').should('have.length.greaterThan', 0);
  });
});

describe('Payments Page', () => {
  beforeEach(() => {
    cy.visit(`${FRONTEND_URL}/payments`);
  });

  it('should display the Payments header', () => {
    cy.get('h2').contains('Payments').should('be.visible');
  });

  it('should have all input fields visible', () => {
    cy.get('input[name="amount"]').should('be.visible');
    cy.get('input[name="cardNumber"]').should('be.visible');
    cy.get('input[name="expirationDate"]').should('be.visible');
  });

  it('should allow entering payment details', () => {
    cy.get('input[name="amount"]').type('100');
    cy.get('input[name="cardNumber"]').type('1234567812345678');
    cy.get('input[name="expirationDate"]').type('12/25');
  });

  it('should submit the form and display a success message on valid input', () => {
    cy.intercept('POST', 'http://localhost:8080/payments', {
      statusCode: 200,
      body: { message: 'Payment Successful!' },
    }).as('postPayment');

    cy.get('input[name="amount"]').type('100');
    cy.get('input[name="cardNumber"]').type('1234567812345678');
    cy.get('input[name="expirationDate"]').type('12/25');
    cy.get('button[type="submit"]').click();

    cy.wait('@postPayment');
    cy.contains('Payment Successful!').should('be.visible');
  });

  it('should display an error message on failed payment', () => {
    cy.intercept('POST', 'http://localhost:8080/payments', {
      statusCode: 500,
      body: { message: 'Payment Failed!' },
    }).as('postPayment');

    cy.get('input[name="amount"]').type('100');
    cy.get('input[name="cardNumber"]').type('1234567812345678');
    cy.get('input[name="expirationDate"]').type('12/25');
    cy.get('button[type="submit"]').click();

    cy.wait('@postPayment');
    cy.contains('Payment Failed!').should('be.visible');
  });

  it('should not submit the form if required fields are empty', () => {
    cy.get('button[type="submit"]').click();
    cy.contains('Payment form submitted').should('not.exist');
  });
});
