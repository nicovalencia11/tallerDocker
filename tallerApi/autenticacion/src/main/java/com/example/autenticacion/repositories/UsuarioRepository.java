package com.example.autenticacion.repositories;

import com.example.autenticacion.entities.Usuario;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface UsuarioRepository extends JpaRepository<Usuario, Integer> {

    /**
     * Login con correo electronico y contraseña
     * @param nombreUsuario
     * @param password
     * @return
     */
    @Query("select u from Usuario u where u.nombreUsuario = :nombreUsuario and u.password = :password")
    Usuario login(String nombreUsuario, String password);

    /**
     * Login por medio del nombre de usuario y contraseña
     * @param usuario
     * @param pass
     * @return
     */
    Usuario findByNombreUsuarioAndPassword(String usuario, String pass);

    Usuario findByCodigo(int codigo);

    /**
     * buscar usuario por correo
     * @param correo
     * @return
     */
    Usuario findByCorreo(String correo);

    List<Usuario> findAll();
}
