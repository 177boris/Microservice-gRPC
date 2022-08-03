import json 
from flask import Blueprint, request, jsonify, session, render_template

from todo.client import TodoClient
from middleware.middleware import token_required

todo_blueprint = Blueprint('todo', __name__)

@todo_blueprint.route('/get', methods=['GET'])
@token_required
def get_todos():
    try:
        user_data = session['user_data']
        client = TodoClient()
        result = client.get_todos(user_id=user_data.get('ID'))
        data = json.loads(result.data)
        return render_template('todo/list.html', todos=data)
    except Exception as error:
        return render_template('todo/list.html', error=str(error))


@todo_blueprint.route('/create', methods=['GET', 'POST'])
@token_required
def create_todo():
    if request.method == 'POST':
        try:
            title = request.form['title']
            description = request.form['description']
            user_data = session['user_data']
            if title == '' or description == '':
                raise ValueError('Please fill the title and description')
            client = TodoClient()
            result = client.create_todo(title=request.form['title'],
                                        description=description,
                                        user_id=user_data.get('ID'))
            return redirect(url_for('todo_blueprint.get_todos'))
        except Exception as error:
            return render_template('todo/create.html', error=str(error))
    return render_template('todo/create.html')
    