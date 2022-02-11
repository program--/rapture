#' Print a sparse rapture grid
show <- function(object) {
    n  <- attr(object, "n")
    bb <- attr(object, "bounds")
    cat("# Rapture Sparse Matrix\n")
    cat(sprintf("# Dims:   %d \u00d7 %d", n, n), "\n")
    cat(sprintf("# Bounds: [%02f, %02f, %02f, %02f]",
                bb$xmin, bb$ymin, bb$xmax, bb$ymax), "\n")

    rn <- nchar(max(head(object@i)))
    cn <- nchar(max(head(object@j)))
    vn <- nchar(max(head(object@x)))

    if (rn < 3) rn <- 3
    if (cn < 3) cn <- 3
    if (vn < 5) vn <- 5

    cat(sprintf(
        paste0("%-3s %", rn, "s %", cn, "s %", vn, "s\n"),
        "",
        "Row",
        "Col",
        "Count"
    ))

    for (k in 1:10) {
        cat(sprintf(
            paste0("%2s  %", rn, "s %", cn, "s %", vn, "s\n"),
            paste0(k),
            object@i[k],
            object@j[k],
            object@x[k]
        ))
    }

    cat(
        "# \u2026 with",
        formatC(
            length(object@i) - 10,
            format = "d",
            big.mark = ","
        ),
        "more entries\n"
    )
}

setClass("rapture_sparse")
setMethod("show", signature("rapture_sparse"), show)