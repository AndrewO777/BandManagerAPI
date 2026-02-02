CREATE DATABASE IF NOT EXISTS BandManager;
USE BandManager;

CREATE TABLE Users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    -- Will add these later, auth not implemented yet
    -- email VARCHAR(100) NOT NULL UNIQUE,
    -- cognito_user_id VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE Bands (
    band_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE UserBand (
    user_id INT NOT NULL,
    band_id INT NOT NULL,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, band_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (band_id) REFERENCES Bands(band_id) ON DELETE CASCADE
);

CREATE TABLE Songs (
    song_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    artist VARCHAR(100),
    key VARCHAR(20),
    time_signature VARCHAR(10),
    bpm DECIMAL(6,1),
    genre VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Setlists (
    setlist_id INT AUTO_INCREMENT PRIMARY KEY,
    band_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    is_template BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (band_id) REFERENCES Bands(band_id) ON DELETE CASCADE
);

CREATE TABLE SetlistSongs (
    setlist_id INT NOT NULL,
    song_id INT NOT NULL,
    position INT NOT NULL,
    PRIMARY KEY (setlist_id, song_id),
    FOREIGN KEY (setlist_id) REFERENCES Setlists(setlist_id) ON DELETE CASCADE,
    FOREIGN KEY (song_id) REFERENCES Songs(song_id) ON DELETE CASCADE
);

CREATE TABLE UserPracticeLogs (
    log_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    song_id INT NOT NULL,
    difficulty_rating TINYINT CHECK (difficulty_rating BETWEEN 1 AND 10),
    last_practiced_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (song_id) REFERENCES Songs(song_id) ON DELETE CASCADE
);

CREATE TABLE BandPracticeLogs (
    log_id INT AUTO_INCREMENT PRIMARY KEY,
    band_id INT NOT NULL,
    song_id INT NOT NULL,
    last_practiced_at TIMESTAMP,
    FOREIGN KEY (band_id) REFERENCES Bands(band_id) ON DELETE CASCADE,
    FOREIGN KEY (song_id) REFERENCES Songs(song_id) ON DELETE CASCADE,
    UNIQUE KEY (band_id, song_id)
);

CREATE INDEX idx_userband_band_id ON UserBand(band_id);
CREATE INDEX idx_setlists_band_id ON Setlists(band_id);
CREATE INDEX idx_setlistsongs_song_id ON SetlistSongs(song_id);
CREATE INDEX idx_userpracticelogs_song_id ON UserPracticeLogs(song_id);
CREATE INDEX idx_bandpracticelogs_song_id ON BandPracticeLogs(song_id);
