@SIGNUP
Feature: Actualizar  un usuario nuevo

  Background: Endpoint y datos para actualizar de usuario
    * url "http://localhost:8081/api/v1/usuario/"
    * request {"codigo": 0, "nombre": "string", "correo": "string", "nombreUsuario": "string", "token": "string", "password": "string"}

  @CP1
  Scenario Outline: Validar que al enviar los datos correspondientes permita actualizar el usuario
    When method put
    Then status 200
    Examples:
      | codigo  |nombre                   | correo | nombreUsuario       | password | token              |
      | 1       | Nicolas valencia        | n@gmail.com | nicovalencia11      | 123456   | hfghfghfghfgh |

  @CP2
  Scenario Outline: Validar que no permita actualizar el usuario con campos obligatorios vacios
    When method put
    Then status 500
    Examples:
      | codigo  |nombre                   | correo | nombreUsuario       | password | token              |
      |        |                         | n@gmail.com | nicovalencia11      | 123456   | hfghfghfghfgh |

  @CP4
  Scenario Outline: Validar que no permita guardar al exceder el tamanio permitido al campo nombre
    When method put
    Then status 500
    And print  'La longitud máxima permitida es 45'
    Examples:
      | codigo  |nombre                   | correo | nombreUsuario       | password | token              |
      | 1       | Nicolas valenciaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa        | n@gmail.com | nicovalencia11      | 123456   | hfghfghfghfgh |

