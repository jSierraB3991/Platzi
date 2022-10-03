package org.jSierra3991.movies.service;

import org.jSierra3991.movies.models.Movie;
import org.jSierra3991.movies.repository.MovieRepository;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnitRunner;


import java.util.List;

import static org.jSierra3991.movies.models.Genre.ACTION;
import static org.jSierra3991.movies.models.Genre.COMEDY;
import static org.jSierra3991.movies.models.Genre.HORROR;
import static org.jSierra3991.movies.models.Genre.THRILLER;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;
import static org.mockito.Mockito.when;

@RunWith(MockitoJUnitRunner.class)
public class MovieServiceTest {

    @InjectMocks
    private MovieService service;
    @Mock
    private MovieRepository repository;

    @Test
    public void return_movies_by_genre() {
        when(repository.findAll())
                .thenReturn(dbMovies());

        var list = service.findByGenre(COMEDY);
        assertFalse(list.isEmpty());
        assertTrue(list.stream().allMatch(m -> m.getGenre().equals(COMEDY)));
        assertEquals(2, list.size());
    }

    @Test
    public void return_movies_by_duration() {
        when(repository.findAll())
                .thenReturn(dbMovies());

        var list = service.minusToMinutes(149);
        assertFalse(list.isEmpty());
        assertTrue(list.stream().allMatch(m -> m.getMinutes() <= 149));
        assertEquals(3, list.size());
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
}