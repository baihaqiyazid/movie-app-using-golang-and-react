// THIRD PARTY
import { Row, Col } from 'react-bootstrap';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'

// COMPONENTS
import { Menu } from './components/Menu';

// PAGES
import { Home } from './pages/Home';
import { Movies } from './pages/Movies';
import { Genres } from './pages/Genres';
import { Admin } from './pages/Admin';
import { MovieDetail } from './pages/Movies/MovieDetail';
import MovieCreate from './pages/Admin/MovieCreate';
import ShowMovies from './pages/Genres/ShowMovies';

//STYLE
import './App.css';

function App() {
  return (
    <Router>
      <div className="mt-4" style={{ margin: "5rem" }}>
        <h1>Movie List</h1>
        <hr className='mb-4' />
        <Row>
          <Col md={2}>
            <Menu />
          </Col>
          <Col>
            <Routes>
              {/* HOME */}
              <Route path='/' element={<Home />} />
              {/* MOVIES */}
              <Route path='/movies' element={<Movies />} />
              <Route exact path='/movies/:id' element={<MovieDetail />} />
              
              {/* GENRES */}
              <Route path='/genres' element={<Genres />} />
              <Route exact path='/genres/:id' element={<ShowMovies />} />
              {/* ADMIN */}
              <Route exact path='/admin' element={<Admin />} />
              <Route path='/admin/movies/create' element={<MovieCreate />} />
              <Route exact path='/admin/movies/:id/edit' element={<MovieCreate />} />
            </Routes>
          </Col>
        </Row>
      </div>
    </Router>

  );
}

export default App;
