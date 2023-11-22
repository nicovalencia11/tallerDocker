@SIGNUP
Feature: Crear  un usuario nuevo

  Background: Endpoint y datos para registro de usuario
    * url "http://localhost:8081/api/v1/usuario/"
    * request { "nombre": "#(nombre)", "nombreUsuario": "#(nombreUsuario)", "password": "#(password)", "correo": "#(correo)" }

  @CP1
  Scenario Outline: Validar que al Ingresar los datos correspondientes permita crear el nuevo usuario
    When method post
    Then status 200
    And match response == 'Registro Exitoso'
    Examples:
      | nombre   | nombreUsuario | password | correo              |
      | nicotrin | nicotrin      | 123456   | nicotrin@correo.com |

  @CP2
  Scenario Outline: Validar que no permita registrar el usuario con campos obligatorios vacios
    When method post
    Then status 500
    Examples:
      | nombre | nombreUsuario | password | correo          |
      |        |               | 123456   | july@correo.com |

  @CP3
  Scenario Outline: Validar que no permita registrar el usuario con resgitros duplicados
    When method post
    Then status 500
    Examples:
      | nombre | nombreUsuario | password | correo           |
      | july   | julys         | 123456   | nicotrin@correo.com |

  @CP4
  Scenario Outline: Validar que no permita guardar al exceder el tamanio permitido al campo nombre
    When method post
    Then status 500
    And print  'La longitud máxima permitida es 45'
    Examples:
      | nombre                                                                                | nombreUsuario | password | correo           |
      | juliiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii | julys         | 123456   | nuevo@correo.com |



  @CP5
  Scenario Outline: Validar que no permita guardar al exceder el tamanio permitido al campo nombre_usuario
    When method post
    Then status 500
    And print  'La longitud máxima permitida es 45'
    Examples:
      | nombre | nombreUsuario                                                                         | password | correo           |
      | july   | juliiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii | 123456   | nuevo@correo.com |