Feature: registro

  Scenario: quiero registrar un usuario
    Given debe estar el micro activo
    When se realiza la peticion post con body json
    Then veo un mensaje de confirmación
