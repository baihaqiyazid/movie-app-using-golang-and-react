import React from 'react'
import { Card, Col, ListGroup, Nav, NavItem, NavLink } from 'react-bootstrap';
import { Link } from 'react-router-dom';

export const Menu = () => {
  return (

    <Card className='mb-3'>
      <ListGroup variant="flush">
        <ListGroup.Item> <Link to="/">Home</Link></ListGroup.Item>
        <ListGroup.Item> <Link to="/movies">Movies</Link></ListGroup.Item>
        <ListGroup.Item> <Link to="/genres">Genres</Link></ListGroup.Item>
        <ListGroup.Item> <Link to="/admin">Admin</Link></ListGroup.Item>
      </ListGroup>
    </Card>
  )
}
