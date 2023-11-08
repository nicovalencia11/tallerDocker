Feature: get user on regres

  Scenario: Get a user
    Given url "http://localhost:8080/api/v1/usuario/"
    When method get
    Then status 200
