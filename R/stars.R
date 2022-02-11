#' Coerce a `Matrix::sparseMatrix` to `stars`.
#' @param x A matrix that inherits `Matrix::sparseMatrix`
#' @rdname st_as_stars
#' @export
st_as_stars.rapture_sparse <- function(.x, ...) {
    class(.x) <- "dgTMatrix"
    obj <- stars::st_as_stars(as.matrix(.x))
    obj <- stars::st_set_bbox(obj, attr(.x, "bounds"))
    obj <- sf::st_set_crs(obj, 4326)
    obj <- sf::st_transform(obj, 3857)
    obj[[1]][obj[[1]] == 0] <- NA
    attr(obj, "bounds") <- attr(.x, "bounds")
    attr(obj, "n")      <- attr(.x, "n")
    return(obj)
}

#' @export
st_as_stars.sparseMatrix <- function(.x, ...) {
    obj <- stars::st_as_stars(as.matrix(.x))
    obj <- stars::st_set_bbox(obj, attr(.x, "bounds"))
    obj <- sf::st_set_crs(obj, 4326)
    obj <- sf::st_transform(obj, 3857)
    obj[[1]][obj[[1]] == 0] <- NA
    attr(obj, "bounds") <- attr(.x, "bounds")
    attr(obj, "n")      <- attr(.x, "n")
    return(obj)
}