-- 存储图片信息
CREATE TABLE images (
    image_id INT PRIMARY KEY AUTO_INCREMENT,#由gorm.Model取代
    image_url VARCHAR(255),
    category ENUM('General','Anime','People'),
    purity ENUM('SFW','Sketchy','NSFW'),
    user_id INT
);

-- 存储标签信息
CREATE TABLE tags (
    tag_id INT PRIMARY KEY AUTO_INCREMENT,#由gorm.Model取代
    tag_name VARCHAR(50)
);

-- 关联表，存储图片与标签之间的关系
CREATE TABLE image_tag (
    image_id INT,
    tag_id INT,
    PRIMARY KEY (image_id, tag_id)
);
