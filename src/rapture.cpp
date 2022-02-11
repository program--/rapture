#include <cpp11.hpp>
#include <unordered_set>
#include "grid.h"

// STL
using std::pair;
using std::size_t;
using std::unordered_set;
using std::vector;

// cpp11
namespace writable = cpp11::writable;
using namespace cpp11::literals;
using cpp11::as_cpp;
using cpp11::by_column;
using cpp11::data_frame;
using cpp11::doubles;
using cpp11::integers;
using cpp11::integers_matrix;

doubles CPP_bounds(doubles x, doubles y) {
    auto fn = cpp11::package("rapture")[".bounds"];
    return as_doubles(fn(x, y));
}

integers_matrix<by_column> empty_matrix(size_t rows, size_t cols) {
    writable::integers_matrix<by_column> mat(rows, cols);
    for (size_t i = 0; i < cols; ++i) {
        for (size_t j = 0; j < rows; ++j) {
            mat[i][j] = 0;
        }
    }
    return mat;
}

[[cpp11::register]]
integers_matrix<by_column> CPP_empty_grid(size_t n) { return empty_matrix(n, n); }

[[cpp11::register]]
integers_matrix<by_column> CPP_dense_grid(doubles x, doubles y, size_t n) {
    size_t len = x.size();

    // Index X and Y coordinates to grid
    vector<size_t> gridx = rapture::index(n, as_cpp<vector<double>>(x)),
                   gridy = rapture::index(n, as_cpp<vector<double>>(y));

    writable::integers_matrix<by_column> grid = CPP_empty_grid(n);
    for (size_t i = 0; i < len; ++i) {
        grid[gridy[i]][gridx[i]] += 1;
    }

    return grid;
}

// https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames#C.2FC.2B.2B
size_t long2tilex(double lon, size_t z) { 
	return static_cast<size_t>(floor((lon + 180.0) / 360.0 * (1 << z))); 
}

// https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames#C.2FC.2B.2B
size_t lat2tiley(double lat, size_t z) { 
    double latrad = lat * M_PI / 180.0;
	return static_cast<size_t>(floor((1.0 - asinh(tan(latrad)) / M_PI) / 2.0 * (1 << z))); 
}

struct hash {
    size_t operator()(const pair<size_t, size_t> &x) const {
        return x.first ^ x.second;
    }
};

[[cpp11::register]]
data_frame CPP_wgs_to_tile(doubles x, doubles y, size_t z) {
    size_t len = x.size();
    unordered_set<pair<size_t, size_t>, hash> tileset;
    for (size_t i = 0; i < len; ++i) {
        tileset.insert({ long2tilex(x[i], z), lat2tiley(y[i], z) });
    }

    size_t ulen = tileset.size();
    vector<size_t> tiles_x,
                   tiles_y;
    for (auto it = tileset.begin(); it != tileset.end(); ++it) {
        tiles_x.push_back(it->first);
        tiles_y.push_back(it->second);
    }

    return writable::data_frame({
        "tileX"_nm = tiles_x,
        "tileY"_nm = tiles_y
    });
}