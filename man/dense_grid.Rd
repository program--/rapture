% Generated by roxygen2: do not edit by hand
% Please edit documentation in R/grid.R
\name{dense_grid}
\alias{dense_grid}
\title{Get a dense matrix from X/Y coordinates}
\usage{
dense_grid(x, y, n = 2^10)
}
\arguments{
\item{x}{Vector of X coordinates}

\item{y}{Vector of Y coordinates}

\item{n}{Grid dimensions, \strong{must be a power of 2}.}
}
\value{
Class of \code{dgCMatrix}
}
\description{
Get a dense matrix from X/Y coordinates
}
