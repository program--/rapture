#' Get a dense matrix from X/Y coordinates
#' @param x Vector of X coordinates
#' @param y Vector of Y coordinates
#' @param n Grid dimensions, **must be a power of 2**.
#' @return Class of `dgCMatrix`
#' @importFrom methods as
#' @export
dense_grid <- function(x, y, n = 2 ^ 10) {
    mat <- methods::as(CPP_dense_grid(x, y, as.integer(n)), "dgCMatrix")
    attr(mat, "bounds") <- .bounds(x, y)
    attr(mat, "n")      <- n
    attr(mat, "class")  <- c("rapture_dense", class(mat))
    return(mat)
}

#' Get a sparse triples matrix from X/Y coordinates
#' @param x A vector of X coordinates, `data.frame`,
#'          `matrix`, or `rapture_dense` object.
#' @param ... Unused.
#' @param y A vector of Y coordinates.
#' @param n Grid dimensions, **must be a power of 2**.
#' @param coords Columns containing X/Y coordinates, respectively.
#' @return Class of `dgTMatrix`
#' @rdname sparse_grid
#' @importFrom methods as
#' @export
sparse_grid <- function(x, ...) {
    UseMethod("sparse_grid")
}

#' @export
sparse_grid.rapture_dense <- function(x, ...) {
    mat <- methods::as(x, "dgTMatrix")
    attr(mat, "bounds") <- attr(x, "bounds")
    attr(mat, "n")      <- attr(x, "n")
    attr(mat, "class")  <- c("rapture_sparse", class(mat))
    return(mat)
}

#' @export
sparse_grid.tbl <- function(x, ..., n = 2 ^ 10, coords = 1:2) {
    x <- as.data.frame(x)
    .Class <- "data.frame"
    NextMethod("sparse_grid", x)
}

#' @export
sparse_grid.data.frame <- function(x, ..., n = 2 ^ 10, coords = 1:2) {
    y <- x[[coords[2]]]
    x <- x[[coords[1]]]
    .Class <- "numeric"
    NextMethod("sparse_grid", x, y = y, n = n)
}

#' @export
sparse_grid.matrix <- function(x, ..., n = 2 ^ 10, coords = 1:2) {
    y <- x[, coords[2]]
    x <- x[, coords[1]]
    .Class <- "numeric"
    NextMethod("sparse_grid", x, y = y, n = n)
}

#' @export
sparse_grid.integer <- function(x, y, ..., n = 2 ^ 10) {
    x <- as.numeric(x)
    y <- as.numeric(y)
    .Class <- "numeric"
    NextMethod("sparse_grid", x, y = y, n = n)
}

#' @export
sparse_grid.double <- function(x, y, ..., n = 2 ^ 10) {
    x <- as.numeric(x)
    y <- as.numeric(y)
    .Class <- "numeric"
    NextMethod("sparse_grid", x, y = y, n = n)
}

#' @export
sparse_grid.numeric <- function(x, y, ..., n = 2 ^ 10) {
    mat <- as(CPP_dense_grid(x, y, as.integer(n)), "dgTMatrix")
    attr(mat, "bounds") <- .bounds(x, y)
    attr(mat, "n")      <- n
    attr(mat, "class")  <- c("rapture_sparse", class(mat))
    return(mat)
}