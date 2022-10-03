package org.jSierra3991.movies.repository.impl;

import lombok.RequiredArgsConstructor;
import org.jSierra3991.movies.models.Genre;
import org.jSierra3991.movies.models.Movie;
import org.jSierra3991.movies.repository.MovieRepository;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;

import java.util.List;
import java.util.stream.Collectors;

@RequiredArgsConstructor
public class MovieRepositoryImpl implements MovieRepository {

    private final JdbcTemplate jdbcTemplate;

    @Override
    public Movie findById(Integer id) {
        RowMapper<Movie> rowMapper = getMovieRowMapper();
        var list = jdbcTemplate.query("SELECT * FROM movie WHERE id = ?",  rowMapper, id);
        return list.isEmpty() ? null : list.get(0);
    }

    @Override
    public List<Movie> findAll() {
        RowMapper<Movie> rowMapper = getMovieRowMapper();
        return jdbcTemplate.query("SELECT * FROM movie", rowMapper);
    }

    private static RowMapper<Movie> getMovieRowMapper() {
        return (rs, rowNum) -> Movie.builder()
                .id(rs.getInt("id"))
                .name(rs.getString("name"))
                .minutes(rs.getInt("minutes"))
                .director(rs.getString("director"))
                .genre(Genre.valueOf(rs.getString("genre")))
                .build();
    }

    @Override
    public List<Movie> findByName(String name) {
        var list = findAll();
        return list.stream()
                .filter(m -> m.getName().toLowerCase().contains(name.toLowerCase()))
                .collect(Collectors.toList());
    }

    @Override
    public List<Movie> findByDirector(String director) {
        var list = findAll();
        return list.stream()
                .filter(m -> m.getDirector().toLowerCase().contains(director.toLowerCase()))
                .collect(Collectors.toList());
    }

    @Override
    public void saveOrUpdate(Movie movie) {
        Object[] args = { movie.getName(), movie.getMinutes(), movie.getGenre().name(), movie.getDirector()};
        jdbcTemplate.update("INSERT INTO movie (name, minutes, genre, director) VALUES(?, ?, ?, ?);", args);
    }
}
