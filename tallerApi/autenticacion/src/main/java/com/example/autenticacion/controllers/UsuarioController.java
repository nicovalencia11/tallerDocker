package com.example.autenticacion.controllers;

import com.example.autenticacion.entities.Usuario;
import com.example.autenticacion.services.implementation.UsuarioServiceImpl;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1/usuario")
public class UsuarioController {

    @Autowired
    private UsuarioServiceImpl usuarioService;

    @Autowired
    private RabbitTemplate rabbitTemplate;

    private final String exchangeName = "rootExchange";
    private final String routingKey = "tuRoutingKey";

    @PostMapping("/")
    public ResponseEntity<String> registrar(@RequestBody Usuario usuario) {
        try {
            usuarioService.registrarUsuario(usuario);
            // Enviar mensaje a RabbitMQ
            rabbitTemplate.convertAndSend(exchangeName, routingKey, "Exito Usuario registrado: " + usuario.getNombreUsuario());
            return ResponseEntity.status(HttpStatus.OK).body("Registro Exitoso");
        } catch (Exception e) {
            rabbitTemplate.convertAndSend(exchangeName, routingKey, "Error al registrar el usuario: " + usuario.getNombreUsuario());
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Error al registrar el usuario: " + e.getMessage());
        }
    }

    @PutMapping("/")
    public ResponseEntity<String> actualizar(@RequestHeader("token") String token, @RequestBody Usuario usuario) throws Exception {
        Usuario usuarioActualizado = usuarioService.buscarUsuario(usuario);
        usuarioActualizado.setPassword(usuario.getPassword());
        if (token.equals(usuarioActualizado.getToken())) {
            usuarioService.actualizarUsuario(usuarioActualizado);
            // Enviar mensaje a RabbitMQ
            rabbitTemplate.convertAndSend(exchangeName, routingKey, "Exito Usuario actualizado: " + usuario.getNombreUsuario());
            return ResponseEntity.status(HttpStatus.OK).body("Actualización Exitosa");
        } else {
            rabbitTemplate.convertAndSend(exchangeName, routingKey, "Error al Actualizar Usuario: " + usuario.getNombreUsuario());
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("No tienes permisos para actualizar este usuario");
        }
    }

    @GetMapping("/")
    public Page<Usuario> listar(@RequestParam(name = "pagina", defaultValue = "0") int pagina,
                                @RequestParam(name = "tamano", defaultValue = "10") int tamano) throws Exception {
        Page<Usuario> usuarios = usuarioService.listarUsuarios(pagina, tamano);
        // Enviar mensaje a RabbitMQ
        rabbitTemplate.convertAndSend(exchangeName, routingKey, "Exito Listado de usuarios solicitado");
        return usuarios;
    }

    @PostMapping("/recuperarPassword")
    public ResponseEntity<String> recuperarPassword(@RequestBody String correo) throws Exception {
        Usuario usuario = usuarioService.recuperarPassword(correo);
        // Enviar mensaje a RabbitMQ
        rabbitTemplate.convertAndSend(exchangeName, routingKey, "Exito Recuperación de contraseña para: " + correo);
        return ResponseEntity.status(HttpStatus.OK).body("Tu clave es: " + usuario.getPassword());
    }
}
