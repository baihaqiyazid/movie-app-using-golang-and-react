import React, { useEffect, useState } from 'react'

// THIRD PARTY
import { Button, Col, FloatingLabel, Form, FormLabel, Row } from 'react-bootstrap'
import { useForm } from 'react-hook-form'
import { Link, useNavigate, useParams } from 'react-router-dom'
import Swal from 'sweetalert2'
import axios from 'axios'

const MovieForm = () => {
    const [genres, setGenres] = useState([])
    const navigate = useNavigate()
    const { register, handleSubmit, setValue, formState: { errors } } = useForm()

    const { id } = useParams()
    const isCreateMovie = !id
    const field = [
        "id", "title", "description", "release_date", "runtime", "rating", "mpaa_rating"
    ]

    const fetchMovies = async (id) => {
        const result = await axios(`http://localhost:4001/movies/${id}`)
        result.data.data.id = result.data.data.id.toString()
        result.data.data.runtime = result.data.data.runtime.toString()
        result.data.data.rating = result.data.data.rating.toString()
        result.data.data.release_date = new Date(result.data.data.release_date).toISOString().split('T')[0];
        field.forEach((field) => setValue(field, result.data.data[field]))
    }
    const fetchGenres = async () => {
        const result = await axios(`http://localhost:4001/genres`)
        await setGenres(result.data.data)
    }

    useEffect(() => {
        if (!isCreateMovie) {
            fetchMovies(id)
        }
        fetchGenres()
    }, [isCreateMovie])

    const onSubmit = async (data) => {

        if (!isCreateMovie) {
            const result = await axios.post(
                `http://localhost:4001/admin/movies/edit`,
                JSON.stringify(data)
            );
            console.log(result.data);
            navigate("/admin")
            Swal.fire(
                'Success!',
                'Data has been updated',
                'success'
            )
        } else {
            const result = await axios.post(
                "http://localhost:4001/admin/movies/create",
                JSON.stringify(data)
            );
            console.log(result.data);
            navigate("/admin")
            Swal.fire(
                'Success!',
                'Data has been created',
                'success'
            )
        }
    };
    return (
        <>
            <Form onSubmit={handleSubmit(onSubmit)} action="/admin">
                {/* TITLE */}
                <FloatingLabel
                    label="Title"
                    className="mb-3"
                >
                    <Form.Control type="text" placeholder="name@example.com"
                        id="title"
                        name="title"
                        {...register("title", { required: true, maxLength: 50 })}
                    />
                    {errors?.title?.type === "required" && (
                        <small className="text-danger mt-1">
                            *This field is required
                        </small>
                    )}
                    {errors?.title?.type === "maxLength" && (
                        <small className="text-danger mt-1">
                            *Title cannot exceed 50 characters
                        </small>
                    )}
                </FloatingLabel>

                {/* GENRE */}
                <FormLabel>Genre</FormLabel>
                <div style={{ height: '150px' }} className="overflow-auto border p-2">
                    {genres.map((genre, index) => (
                        <div key={index} className="mb-1">
                            <Form.Check
                                inline
                                label={genre.name}
                                name="genre_id"
                                id="genre_id"
                                {...register("genre_id", { required: true })}
                                type="checkbox"
                                value={genre.id}
                            />
                        </div>
                    ))}
                </div>
                {errors?.genre_id?.type === "required" && (
                    <small className="text-danger mt-1">
                        *This field is required
                    </small>
                )}

                {/* Release Date */}
                <FloatingLabel
                    label="Release Date"
                    className="mb-3 mt-3"
                >
                    <Form.Control
                        type="date"
                        id="release_date"
                        name="release_date"
                        {...register("release_date", { required: true })}
                    />
                    {errors?.release_date?.type === "required" && (
                        <small className="text-danger mt-1">
                            *This field is required
                        </small>
                    )}
                </FloatingLabel>


                <Row>
                    <Col>
                        {/* Runtime */}
                        <FloatingLabel
                            label="Runtime"
                            className="mb-3"
                        >
                            <Form.Control type="number" placeholder="name@example.com"
                                id="runtime"
                                name="runtime"
                                min={0}
                                {...register("runtime", { required: true })}
                            />
                            {errors?.runtime?.type === "required" && (
                                <small className="text-danger mt-1">
                                    *This field is required
                                </small>
                            )}
                        </FloatingLabel>
                    </Col>

                    <Col>
                        {/* MPAA RATING */}
                        <FloatingLabel
                            label="MPAA Rating"
                            className="mb-3"
                        >
                            <Form.Select id="mpaa_rating" name="mpaa_rating" defaultValue="G" {...register("mpaa_rating", { required: true })}>
                                <option value={'G'}>G</option>
                                <option value={'PG'}>PG</option>
                                <option value={'PG-13'}>PG-13</option>
                                <option value={'R'}>R</option>
                                <option value={'NC-17'}>NC-17</option>
                            </Form.Select>
                            {errors?.mpaa_rating?.type === "required" && (
                                <small className="text-danger mt-1">
                                    *This field is required
                                </small>
                            )}
                        </FloatingLabel>
                    </Col>

                    <Col>
                        {/* rating */}
                        <FloatingLabel
                            label="Rating"
                            className="mb-3"
                        >
                            <Form.Control
                                type="number"
                                placeholder=""
                                id="rating"
                                name="rating"
                                className="d-flex"
                                min={1}
                                max={5}
                                defaultValue={5}
                                {...register("rating", { required: true, max: 5 })}
                            />
                            {errors?.rating?.type === "required" && (
                                <small className="text-danger mt-1">
                                    *This field is required
                                </small>
                            )}
                        </FloatingLabel>
                    </Col>
                </Row>

                <FloatingLabel label="Description">
                    <Form.Control
                        as="textarea"
                        id="description"
                        name="description"
                        placeholder="Leave a description here"
                        style={{ height: '100px' }}
                        {...register("description", { required: true })}
                    />
                    {errors?.description?.type === "required" && (
                        <small className="text-danger mt-1">
                            *This field is required
                        </small>
                    )}
                </FloatingLabel>

                <Button variant="primary" type="submit" className='mt-3'>
                    Submit
                </Button>

            </Form>
        </>
    )
}

export default MovieForm