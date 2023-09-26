package com.example.autenticacion.services.services;

import com.example.autenticacion.entities.Usuario;
import org.springframework.data.domain.Page;

import java.util.List;

/**
 * Interface con los metodos abstractos para el servicio de usuario
 */
public interface UsuarioService {

    /**
     * Metodo que permite registrar un usuario
     * @param usuario
     * @return
     */
    Usuario registrarUsuario(Usuario usuario) throws Exception;

    /**
     * Metodo que permite el login con nombre de usuario y contraseña
     * @param nombreUsuario
     * @param password
     * @return
     */
    Usuario loginUsuario(String nombreUsuario, String password) throws Exception;

    /**
     * Metodo que permite actualizar un usuario
     * @param usuario
     * @return
     */
    Usuario actualizarUsuario(Usuario usuario) throws Exception;

    /**
     * metodo que permite listar todos los usuarios
     * @return
     */
    Page<Usuario> listarUsuarios(int pagina, int tamano);

    /**
     * Metodo que permite recuperar la clave con el correo
     * @param correo
     * @return
     */
    Usuario recuperarPassword(String correo) throws Exception;

    Usuario buscarUsuario(Usuario usuario) throws Exception;

}
