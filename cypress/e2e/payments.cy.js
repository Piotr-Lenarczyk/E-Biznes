const FRONTEND_URL = 'http://172.23.224.32:3000';

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