const add = document.querySelector("#add");
const form = add.querySelector("#add-form");
const inputArchive = add.querySelector("#input-archive");
const lableArchiveImage = add.querySelector("#label-archive-image");
const inputArchiveDescription = add.querySelector("#input-archive-description");

inputArchive.addEventListener("change", (e) => {
  const [file] = inputArchive.files;
  if (file) {
    let ext = file.name.split(".").pop();
    if (ext === "zip") {
      lableArchiveImage.src = "/static/icons/check-mark.png";
      inputArchiveDescription.innerHTML = file.name;
    } else {
      lableArchiveImage.src = "/static/icons/plus.png";
      inputArchiveDescription.innerHTML = "Выберите zip архив.";
      inputArchive.files = null;
    }
  }
});

inputArchive.addEventListener("invalid", function (e) {
  inputArchiveDescription.innerHTML = "Пожалуйста, выберите файл.";
});
