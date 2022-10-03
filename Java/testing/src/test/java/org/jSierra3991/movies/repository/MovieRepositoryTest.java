package org.jSierra3991.movies.repository;

import org.jSierra3991.movies.models.Movie;
import org.jSierra3991.movies.repository.impl.MovieRepositoryImpl;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.springframework.core.io.ClassPathResource;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.datasource.DriverManagerDataSource;
import org.springframework.jdbc.datasource.init.ScriptUtils;

import java.sql.SQLException;
import java.sql.Statement;
import java.util.List;

import static org.jSierra3991.movies.models.Genre.ACTION;
import static org.jSierra3991.movies.models.Genre.COMEDY;
import static org.jSierra3991.movies.models.Genre.DRAMA;
import static org.jSierra3991.movies.models.Genre.HORROR;
import static org.jSierra3991.movies.models.Genre.THRILLER;
import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;

public class MovieRepositoryTest {

    private MovieRepositoryImpl repository;
    private DriverManagerDataSource dataSource;

    @Before
    public void setUp() throws Exception {
        dataSource = new DriverManagerDataSource("jdbc:h2:mem:test;MODE=MYSQL", "sa", "sa");
        ScriptUtils.executeSqlScript(dataSource.getConnection(), new ClassPathResource("database.sql"));
        var template = new JdbcTemplate(dataSource);
        repository = new MovieRepositoryImpl(template);

    }

    @Test
    public void load_all_movie() {
        var result = repository.findAll();
        assertThat(result, is(dbMovies()));
    }

    public List<Movie> dbMovies() {
        return  List.of(Movie.builder().id(1).name("Dark Knight").minutes(150).genre(ACTION).build(),
                Movie.builder().id(2).name("Memento").minutes(150).genre(THRILLER).build(),
                Movie.builder().id(3).name("There Something About Mary").minutes(119).genre(COMEDY).build(),
                Movie.builder().id(4).name("Super 8").minutes(112).genre(THRILLER).build(),
                Movie.builder().id(5).name("Scream").minutes(150).genre(HORROR).build(),
                Movie.builder().id(6).name("Home Alone").minutes(150).genre(COMEDY).build(),
                Movie.builder().id(7).name("Matrix").minutes(136).genre(ACTION).build());
    }

    @Test
    public void find_movie_by_id() {
        var result = repository.findById(3);
        assertThat(result, is(Movie.builder().id(3).name("There Something About Mary").minutes(119).genre(COMEDY).build()));
    }

    @Test
    public void insert_movie_in_database() {
        var movie = Movie.builder().genre(DRAMA).minutes(239).name("Clannad").build();
        repository.saveOrUpdate(movie);

        var result = repository.findById(8);
        movie.setId(8);
        assertThat(result, is(movie));
    }

    @After
    public void tearDown() throws Exception {
        final Statement s = dataSource.getConnection().createStatement();
        s.execute("shutdown");
    }
}