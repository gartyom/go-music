const add = document.querySelector("#add");
const form = add.querySelector("#add-form");
const inputCover = add.querySelector("#input-cover");
const inputArtist = add.querySelector("#input-artist");
const inputRelease = add.querySelector("#input-release");
const inputTrack = add.querySelector("#input-track");
const lableTrackImage = add.querySelector("#label-track-image");
const lableCoverImage = add.querySelector("#label-cover-image");
const inputCoverDescription = add.querySelector("#input-cover-description");
const inputTrackDescription = add.querySelector("#input-track-description");

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
        lableCoverImage.src = "/static/icons/plus.png";
        inputCoverDescription.innerHTML = "Выберите файл (jpeg/jpg)";
      }
      inputCover.files = null;
    }
  }
});

inputTrack.addEventListener("change", (e) => {
  const [file] = inputTrack.files;
  if (file) {
    let ext = file.name.split(".").pop();
    if (ext === "rar") {
      lableTrackImage.src = "/static/icons/check-mark.png";
      inputTrackDescription.innerHTML = file.name;
    } else {
      lableTrackImage.src = "/static/icons/plus.png";
      inputTrackDescription.innerHTML = "Добавьте архив с песнями (rar)";
      inputTrack.files = null;
    }
  }
});

// async function submitAddForm(e) {
//   let data = {
//     "release-image": inputCover.files,
//     "release-tracks": inputTrack.files,
//     artist: inputArtist.value,
//     release: inputRelease.value,
//   };

//   let resopnse = await fetch("/add", {
//     method: "POST",
//     body: JSON.stringify(data),
//     headers: {
//       "Content-type": "application/json; charset=UTF-8",
//     },
//   });

//   let json = resopnse.json();
//   if (json.status === "OK") {
//     form.reset();
//     alert("Добавлено");
//   }
// }
