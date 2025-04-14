const FRONTEND_URL = 'http://172.23.224.32:3000';

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