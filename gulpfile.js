'use strict';
const gulp = require('gulp');
const sass = require('gulp-sass');
const sassSelector = './sass/**/*.scss';
const destPath = './web/';


function css() {
    return gulp.src(sassSelector)
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest(destPath + 'css/'));
}

function watchFiles() {
    gulp.watch(sassSelector, css);
}

const watch = gulp.parallel(watchFiles);

exports.css = css;
exports.watch = watch;