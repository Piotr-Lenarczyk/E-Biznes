const FRONTEND_URL = 'http://172.23.224.32:3000';

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