#' Create a raster capture `tmap`.
#' @param x An object coercable to a `stars` object.
#'          Notably, this should be a `rapture_sparse` object.
#' @param buffer Amount to buffer tiles. See `sf::st_buffer`.
#' @inheritParams basemaps::basemap
#' @param downsample Downsample raster/basemap if it's too big?
#' @param ... Passed to `tmap::tm_raster`.
#' @param basemap_extra_args Named list of arguments passed
#'                           to `basemaps::basemap`.
#' @return A `tmap` object
#' @export
rapture <- function(
    x,
    ...,
    basemap = TRUE,
    buffer = 10,
    map_service = "carto",
    map_type = "voyager_no_labels",
    downsample = FALSE,
    basemap_extra_args = NULL,
    verbose = FALSE
) {
    x   <- stars::st_as_stars(x)
    map <- tmap::tm_shape(x, raster.downsample = downsample) +
           tmap::tm_raster(legend.show = FALSE, ...)

    if (basemap) {
        bb <- sf::st_bbox(sf::st_buffer(
            sf::st_set_crs(sf::st_as_sfc(attr(x, "bounds")), 4326),
            buffer,
            endCapStyle = "SQUARE"
        ))

        args <- c(list(
            ext = bb,
            map_service = map_service,
            map_type = map_type,
            verbose = verbose
        ), basemap_extra_args)

        st_basemap <- do.call(basemaps::basemap_stars, args)

        map <- tmap::tm_shape(st_basemap, raster.downsample = downsample) +
               tmap::tm_rgb() +
               map
    }

    return(map)
}