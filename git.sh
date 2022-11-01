#!/bin/bash

message=""
case $1 in

  # 🎨: art
  art)
    message="Improve structure / format of the code"
    ;;

  # ⚡️: zap
  zap)
    message="Improve performance"
    ;;

  # 🔥: fire
  fire)
    message="Remove code or files"
    ;;

  # 🐛: bug
  bug)
    message="Fix a bug"
    ;;

  # 🚑️: ambulance
  ambulance)
    message="Critical hotfix"
    ;;

  # ✨: sparkles
  sparkles)
    message="Introduce new features"
    ;;

  # 📝: memo
  memo|doc|documentation)
    message="Add or update documentation"
    emoji="memo"
    ;;

  # 🚀: rocket
  rocket)
    message="Deploy stuff"
    ;;

  # 💄: lipstick
  lipstick)
    message="Add or update the UI and style files"
    ;;

  # 🎉: tada
  tada)
    message="Begin a project"
    ;;

  # ✅: white-check-mark
  white_check_mark)
    message="Add, update, or pass tests"
    ;;

  # 🔒️: lock
  lock)
    message="Fix security issues"
    ;;

  # 🔐: closed-lock-with-key
  closed_lock_with_key)
    message="Add or update secrets"
    ;;

  # 🔖: bookmark
  bookmark)
    message="Release / Version tags"
    ;;

  # 🚨: rotating-light
  rotating_light)
    message="Fix compiler / linter warnings"
    ;;

  # 🚧: construction
  construction)
    message="Work in progress"
    ;;

  # 💚: green-heart
  green_heart)
    message="Fix CI Build"
    ;;

  # ⬇️: arrow-down
  arrow_down)
    message="Downgrade dependencies"
    ;;

  # ⬆️: arrow-up
  arrow_up)
    message="Upgrade dependencies"
    ;;

  # 📌: pushpin
  pushpin)
    message="Pin dependencies to specific versions"
    ;;

  # 👷: construction-worker
  construction_worker)
    message="Add or update CI build system"
    ;;

  # 📈: chart-with-upwards-trend
  chart_with_upwards_trend)
    message="Add or update analytics or track code"
    ;;

  # ♻️: recycle
  recycle)
    message="Refactor code"
    ;;

  # ➕: heavy-plus-sign
  heavy_plus_sign)
    message="Add a dependency"
    ;;

  # ➖: heavy-minus-sign
  heavy_minus_sign)
    message="Remove a dependency"
    ;;

  # 🔧: wrench
  wrench)
    message="Add or update configuration files"
    ;;

  # 🔨: hammer
  hammer)
    message="Add or update development scripts"
    ;;

  # 🌐: globe-with-meridians
  globe_with_meridians)
    message="Internationalization and localization"
    ;;

  # ✏️: pencil2
  pencil2)
    message="Fix typos"
    ;;

  # 💩: poop
  poop)
    message="Write bad code that needs to be improved"
    ;;

  # ⏪️: rewind
  rewind)
    message="Revert changes"
    ;;

  # 🔀: twisted-rightwards-arrows
  twisted_rightwards_arrows)
    message="Merge branches"
    ;;

  # 📦️: package
  package)
    message="Add or update compiled files or packages"
    ;;

  # 👽️: alien
  alien)
    message="Update code due to external API changes"
    ;;

  # 🚚: truck
  truck)
    message="Move or rename resources (e.g.: files, paths, routes)"
    ;;

  # 📄: page-facing-up
  page_facing_up)
    message="Add or update license"
    ;;

  # 💥: boom
  boom)
    message="Introduce breaking changes"
    ;;

  # 🍱: bento
  bento)
    message="Add or update assets"
    ;;

  # ♿️: wheelchair
  wheelchair)
    message="Improve accessibility"
    ;;

  # 💡: bulb
  bulb)
    message="Add or update comments in source code"
    ;;

  # 🍻: beers
  beers)
    message="Write code drunkenly"
    ;;

  # 💬: speech-balloon
  speech_balloon)
    message="Add or update text and literals"
    ;;

  # 🗃️: card-file-box
  card_file_box)
    message="Perform database related changes"
    ;;

  # 🔊: loud-sound
  loud_sound)
    message="Add or update logs"
    ;;

  # 🔇: mute
  mute)
    message="Remove logs"
    ;;

  # 👥: busts-in-silhouette
  busts_in_silhouette)
    message="Add or update contributor(s)"
    ;;

  # 🚸: children-crossing
  children_crossing)
    message="Improve user experience / usability"
    ;;

  # 🏗️: building-construction
  building_construction)
    message="Make architectural changes"
    ;;

  # 📱: iphone
  iphone)
    message="Work on responsive design"
    ;;

  # 🤡: clown-face
  clown_face)
    message="Mock things"
    ;;

  # 🥚: egg
  egg)
    message="Add or update an easter egg"
    ;;

  # 🙈: see-no-evil
  see_no_evil)
    message="Add or update a .gitignore file"
    ;;

  # 📸: camera-flash
  camera_flash)
    message="Add or update snapshots"
    ;;

  # ⚗️: alembic
  alembic|experiments|experiment|xp)
    message="Perform experiments"
    emoji="alembic"
    ;;

  # 🔍️: mag
  mag)
    message="Improve SEO"
    ;;

  # 🏷️: label
  label)
    message="Add or update types"
    ;;

  # 🌱: seedling
  seedling)
    message="Add or update seed files"
    ;;

  # 🚩: triangular-flag-on-post
  triangular_flag_on_post)
    message="Add, update, or remove feature flags"
    ;;

  # 🥅: goal-net
  goal_net)
    message="Catch errors"
    ;;

  # 💫: animation
  dizzy)
    message="Add or update animations and transitions"
    ;;

  # 🗑️: wastebasket
  wastebasket)
    message="Deprecate code that needs to be cleaned up"
    ;;

  # 🛂: passport-control
  passport_control)
    message="Work on code related to authorization, roles and permissions"
    ;;

  # 🩹: adhesive-bandage
  adhesive_bandage)
    message="Simple fix for a non-critical issue"
    ;;

  # 🧐: monocle-face
  monocle_face)
    message="Data exploration/inspection"
    ;;

  # ⚰️: coffin
  coffin)
    message="Remove dead code"
    ;;

  # 🧪: test-tube
  test_tube)
    message="Add a failing test"
    ;;

  # 👔: necktie
  necktie)
    message="Add or update business logic $2"
    ;;

  # 🩺: stethoscope
  stethoscope)
    message="Add or update healthcheck"
    ;;

  # 🧱: bricks
  bricks)
    message="Infrastructure related changes"
    ;;

  # 🧑‍💻: technologist
  technologist)
    message="Improve developer experience $2"
    ;;

  # 💸: money-with-wings
  money_with_wings)
    message="Add sponsorships or money related infrastructure"
    ;;

  # 🧵: thread
  thread)
    message="Add or update code related to multithreading or concurrency"
    ;;

  # 🦺: safety-vest
  safety_vest)
    message="Add or update code related to validation"
    ;;

  *)
  message="updated"
  ;;

esac

if [ -z "$emoji" ]
then
    # empty
    if [ -z "$2" ]
    then
        # empty
        git add .; git commit -m ":$1: $message."; git push
    else
        # not empty
        git add .; git commit -m ":$1: $message: $2"; git push
    fi

else
    # not empty
    if [ -z "$2" ]
    then
        # empty
        git add .; git commit -m ":$emoji: $message."; git push
    else
        # not empty
        git add .; git commit -m ":$emoji: $message: $2"; git push
    fi

fi


