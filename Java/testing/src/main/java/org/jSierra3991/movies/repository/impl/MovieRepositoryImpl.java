package org.jSierra3991.movies.repository.impl;

import lombok.RequiredArgsConstructor;
import org.jSierra3991.movies.models.Genre;
import org.jSierra3991.movies.models.Movie;
import org.jSierra3991.movies.repository.MovieRepository;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;

import java.util.List;

@RequiredArgsConstructor
public class MovieRepositoryImpl implements MovieRepository {

    private final JdbcTemplate jdbcTemplate;

    @Override
    public Movie findById(Integer id) {
        return null;
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
                .genre(Genre.valueOf(rs.getString("genre")))
                .build();
    }

    @Override
    public Movie saveOrUpdate(Movie movie) {
        return null;
    }
}
