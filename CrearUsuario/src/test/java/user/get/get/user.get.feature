Feature: obtener todos los usuarios

  Scenario: Get a user
    Given url "http://localhost:8081/api/v1/usuario/"
    When method get
    Then status 200
