const add = document.querySelector("#add");
const form = add.querySelector("#add-form");
const inputCover = add.querySelector("#input-cover");
const inputArtist = add.querySelector("#input-artist");
const inputRelease = add.querySelector("#input-release");
const inputSong = add.querySelector("#input-song");
const lableSongImage = add.querySelector("#label-song-image");
const lableCoverImage = add.querySelector("#label-cover-image");
const inputCoverDescription = add.querySelector("#input-cover-description");
const inputSongDescription = add.querySelector("#input-song-description");

inputCover.addEventListener("change", (e) => {
  const [file] = inputCover.files;
  if (file) {
    let ext = file.name.split(".").pop();
    if (ext === "jpg" || ext === "jpeg") {
      lableCoverImage.src = URL.createObjectURL(file);
      lableCoverImage.classList.remove("w-8", "h-8");
      lableCoverImage.classList.add("w-auto", "h-full");
      inputCoverDescription.innerHTML = file.name;
    } else {
      if (lableCoverImage.classList.contains("w-auto")) {
        lableCoverImage.classList.remove("w-auto", "h-full");
        lableCoverImage.classList.add("w-8", "h-8");
      }
      lableCoverImage.src = "/static/icons/plus.png";
      inputCoverDescription.innerHTML = "Выберите файл (jpeg/jpg)!";
      inputCover.files = null;
    }
  }
});

inputSong.addEventListener("change", (e) => {
  const [file] = inputSong.files;
  if (file) {
    let ext = file.name.split(".").pop();
    if (ext === "rar") {
      lableSongImage.src = "/static/icons/check-mark.png";
      inputSongDescription.innerHTML = file.name;
    } else {
      lableSongImage.src = "/static/icons/plus.png";
      inputSongDescription.innerHTML = "Добавьте архив с песнями (rar)!";
      inputSong.files = null;
    }
  }
});
