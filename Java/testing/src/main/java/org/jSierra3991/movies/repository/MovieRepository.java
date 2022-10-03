package org.jSierra3991.movies.repository;

import org.jSierra3991.movies.models.Movie;

import java.util.List;

public interface MovieRepository {

    Movie findById(Integer id);

    List<Movie> findByName(String name);

    List<Movie> findByDirector(String director);

    List<Movie> findAll();

    void saveOrUpdate(Movie movie);
}