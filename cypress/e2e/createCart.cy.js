const FRONTEND_URL = 'http://172.23.224.32:3000';

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