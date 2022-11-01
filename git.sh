#!/bin/bash

message=""
case $1 in

  # 🎨: art
  art)
    message="Improve structure / format of the code. $2"
    ;;

  # ⚡️: zap
  zap)
    message="Improve performance. $2"
    ;;

  # 🔥: fire
  fire)
    message="Remove code or files. $2"
    ;;

  # 🐛: bug
  bug)
    message="Fix a bug. $2"
    ;;

  # 🚑️: ambulance
  ambulance)
    message="Critical hotfix. $2"
    ;;

  # ✨: sparkles
  sparkles)
    message="Introduce new features. $2"
    ;;

  # 📝: memo
  memo)
    message="Add or update documentation. $2"
    ;;

  # 🚀: rocket
  rocket)
    message="Deploy stuff. $2"
    ;;

  # 💄: lipstick
  lipstick)
    message="Add or update the UI and style files. $2"
    ;;

  # 🎉: tada
  tada)
    message="Begin a project. $2"
    ;;

  # ✅: white-check-mark
  white_check_mark)
    message="Add, update, or pass tests. $2"
    ;;

  # 🔒️: lock
  lock)
    message="Fix security issues. $2"
    ;;

  # 🔐: closed-lock-with-key
  closed_lock_with_key)
    message="Add or update secrets. $2"
    ;;

  # 🔖: bookmark
  bookmark)
    message="Release / Version tags. $2"
    ;;

  # 🚨: rotating-light
  rotating_light)
    message="Fix compiler / linter warnings. $2"
    ;;

  # 🚧: construction
  construction)
    message="Work in progress. $2"
    ;;

  # 💚: green-heart
  green_heart)
    message="Fix CI Build. $2"
    ;;

  # ⬇️: arrow-down
  arrow_down)
    message="Downgrade dependencies. $2"
    ;;

  # ⬆️: arrow-up
  arrow_up)
    message="Upgrade dependencies. $2"
    ;;

  # 📌: pushpin
  pushpin)
    message="Pin dependencies to specific versions. $2"
    ;;

  # 👷: construction-worker
  construction_worker)
    message="Add or update CI build system. $2"
    ;;

  # 📈: chart-with-upwards-trend
  chart_with_upwards_trend)
    message="Add or update analytics or track code. $2"
    ;;

  # ♻️: recycle
  recycle)
    message="Refactor code. $2"
    ;;

  # ➕: heavy-plus-sign
  heavy_plus_sign)
    message="Add a dependency. $2"
    ;;

  # ➖: heavy-minus-sign
  heavy_minus_sign)
    message="Remove a dependency. $2"
    ;;

  # 🔧: wrench
  wrench)
    message="Add or update configuration files. $2"
    ;;

  # 🔨: hammer
  hammer)
    message="Add or update development scripts. $2"
    ;;

  # 🌐: globe-with-meridians
  globe_with_meridians)
    message="Internationalization and localization. $2"
    ;;

  # ✏️: pencil2
  pencil2)
    message="Fix typos. $2"
    ;;

  # 💩: poop
  poop)
    message="Write bad code that needs to be improved. $2"
    ;;

  # ⏪️: rewind
  rewind)
    message="Revert changes. $2"
    ;;

  # 🔀: twisted-rightwards-arrows
  twisted_rightwards_arrows)
    message="Merge branches. $2"
    ;;

  # 📦️: package
  package)
    message="Add or update compiled files or packages. $2"
    ;;

  # 👽️: alien
  alien)
    message="Update code due to external API changes. $2"
    ;;

  # 🚚: truck
  truck)
    message="Move or rename resources (e.g.: files, paths, routes). $2"
    ;;

  # 📄: page-facing-up
  page_facing_up)
    message="Add or update license. $2"
    ;;

  # 💥: boom
  boom)
    message="Introduce breaking changes. $2"
    ;;

  # 🍱: bento
  bento)
    message="Add or update assets. $2"
    ;;

  # ♿️: wheelchair
  wheelchair)
    message="Improve accessibility. $2"
    ;;

  # 💡: bulb
  bulb)
    message="Add or update comments in source code. $2"
    ;;

  # 🍻: beers
  beers)
    message="Write code drunkenly. $2"
    ;;

  # 💬: speech-balloon
  speech_balloon)
    message="Add or update text and literals. $2"
    ;;

  # 🗃️: card-file-box
  card_file_box)
    message="Perform database related changes. $2"
    ;;

  # 🔊: loud-sound
  loud_sound)
    message="Add or update logs. $2"
    ;;

  # 🔇: mute
  mute)
    message="Remove logs. $2"
    ;;

  # 👥: busts-in-silhouette
  busts_in_silhouette)
    message="Add or update contributor(s). $2"
    ;;

  # 🚸: children-crossing
  children_crossing)
    message="Improve user experience / usability. $2"
    ;;

  # 🏗️: building-construction
  building_construction)
    message="Make architectural changes. $2"
    ;;

  # 📱: iphone
  iphone)
    message="Work on responsive design. $2"
    ;;

  # 🤡: clown-face
  clown_face)
    message="Mock things. $2"
    ;;

  # 🥚: egg
  egg)
    message="Add or update an easter egg. $2"
    ;;

  # 🙈: see-no-evil
  see_no_evil)
    message="Add or update a .gitignore file. $2"
    ;;

  # 📸: camera-flash
  camera_flash)
    message="Add or update snapshots. $2"
    ;;

  # ⚗️: alembic
  alembic)
    message="Perform experiments. $2"
    ;;

  # 🔍️: mag
  mag)
    message="Improve SEO. $2"
    ;;

  # 🏷️: label
  label)
    message="Add or update types. $2"
    ;;

  # 🌱: seedling
  seedling)
    message="Add or update seed files. $2"
    ;;

  # 🚩: triangular-flag-on-post
  triangular_flag_on_post)
    message="Add, update, or remove feature flags. $2"
    ;;

  # 🥅: goal-net
  goal_net)
    message="Catch errors. $2"
    ;;

  # 💫: animation
  dizzy)
    message="Add or update animations and transitions. $2"
    ;;

  # 🗑️: wastebasket
  wastebasket)
    message="Deprecate code that needs to be cleaned up. $2"
    ;;

  # 🛂: passport-control
  passport_control)
    message="Work on code related to authorization, roles and permissions. $2"
    ;;

  # 🩹: adhesive-bandage
  adhesive_bandage)
    message="Simple fix for a non-critical issue. $2"
    ;;

  # 🧐: monocle-face
  monocle_face)
    message="Data exploration/inspection. $2"
    ;;

  # ⚰️: coffin
  coffin)
    message="Remove dead code. $2"
    ;;

  # 🧪: test-tube
  test_tube)
    message="Add a failing test. $2"
    ;;

  # 👔: necktie
  necktie)
    message="Add or update business logic $2"
    ;;

  # 🩺: stethoscope
  stethoscope)
    message="Add or update healthcheck. $2"
    ;;

  # 🧱: bricks
  bricks)
    message="Infrastructure related changes. $2"
    ;;

  # 🧑‍💻: technologist
  technologist)
    message="Improve developer experience $2"
    ;;

  # 💸: money-with-wings
  money_with_wings)
    message="Add sponsorships or money related infrastructure. $2"
    ;;

  # 🧵: thread
  thread)
    message="Add or update code related to multithreading or concurrency. $2"
    ;;

  # 🦺: safety-vest
  safety_vest)
    message="Add or update code related to validation. $2"
    ;;

  # 🧪: test-tube
  experiments)
    message="Experiments"
    emoji="test_tube"
    ;;

  *)
  message="updated. $2"
  ;;

esac

if [ -z "$emoji" ]
then
    # empty
    git add .; git commit -m ":$1: $message"; git push
else
    # not empty
    if [ -z "$2" ]
    then
        git add .; git commit -m ":$emoji: $message"; git push
    else
        git add .; git commit -m ":$emoji: $message: $2"; git push
    fi

fi


