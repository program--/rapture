## usethis namespace: start
#' @useDynLib rapture, .registration = TRUE
## usethis namespace: end
NULL

#' Get an empty matrix from X/Y coordinates
#' @param x Vector of X coordinates
#' @param y Vector of Y coordinates
#' @return an `n` by `n` `matrix`.
#' @keywords internal
empty_grid <- function(n) {
    if (n > 2^15 | n < 2^1) stop("`n` must be within the internal [1, 15].")
    mat <- CPP_empty_grid(n)
    attr(mat, "n") <- n
    return(mat)
}

#' Get bounding box from X/Y vectors
#' @param x Vector of X coordinates
#' @param y Vector of Y coordinates
#' @return Vector of class `bbox` with elements
#'         `xmin`, `ymin`, `xmax`, and `ymax`.
#' @keywords internal
.bounds <- function(x, y) {
    bb <- c("xmin" = min(x), "ymin" = min(y),
            "xmax" = max(x), "ymax" = max(y))
    class(bb) <- "bbox"
    return(bb)
}