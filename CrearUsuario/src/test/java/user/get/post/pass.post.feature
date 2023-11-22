@SIGNUP
Feature: Recuperar contyraseña

  Background: Endpoint y datos para registro de usuario
    * url "http://localhost:8080/api/v1/usuario/recuperarPassword"
    * request {"correo": "#(correo)" }

  @CP1
  Scenario Outline: Validar que que con un correo nos permita recuperar las contraseñas
    When method post
    Then status 200
    Examples:
      | correo              |
      | n@gmail.com         |