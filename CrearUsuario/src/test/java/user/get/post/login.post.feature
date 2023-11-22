@SIGNUP
Feature: loggin de un usuario en el sistema

  Background: Endpoint y datos para loggin de usuario
    * url "http://localhost:8081/api/v1/autenticacion/login"
    * request { "nombreUsuario": "#(nombreUsuario)", "password": "#(password)" }

  @CP1
  Scenario Outline: Validar que al Ingresar los datos correspondientes permita el loggin al usuario
    When method post
    Then status 200
    Examples:
      | nombreUsuario | password |
      | nicovalencia11      | 123456   |

  @CP1
  Scenario Outline: Validar que al Ingresar los datos incorrectos no permita el loggin al usuario
    When method post
    Then status 500
    Examples:
      | nombreUsuario | password |
      | nicovalencia11      | 1234567   |