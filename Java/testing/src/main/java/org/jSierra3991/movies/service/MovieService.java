package org.jSierra3991.movies.service;

import lombok.RequiredArgsConstructor;
import org.jSierra3991.movies.models.Genre;
import org.jSierra3991.movies.models.Movie;
import org.jSierra3991.movies.repository.MovieRepository;

import java.util.List;
import java.util.stream.Collectors;

@RequiredArgsConstructor
public class MovieService {

    private final MovieRepository repository;

    public List<Movie> findByGenre(Genre genre) {
        return repository.findAll().stream()
                .filter(m -> m.getGenre().equals(genre))
                .collect(Collectors.toList());
    }
    public List<Movie> minusToMinutes(int minutes) {
        return repository.findAll().stream()
                .filter(m -> m.getMinutes() <= minutes)
                .collect(Collectors.toList());
    }
}